# Filter Service
Welcome to the filter service. The purpose of this service is to validate messages submitted by the client. See below a guide that will help you get started as well as all the features provided by this service.

- Build with Go version 1.18
- Uses [chi router](github.com/go-chi/chi)
- Uses [sqlite](github.com/mattn/go-sqlite3)

## Steps to download
#### Cloning with submodules
>git clone --recurse-submodules https://github.com/ralugr/filter-service
#### If you already cloned the repo without submodules
> git submodule update --init --recursive
#### You can also clone and start the services separately (without using start.sh)
>git clone https://github.com/ralugr/filter-service  
>git clone https://github.com/ralugr/language-service

## Pulling upstream changes from the project remote
By default, the git pull command recursively fetches submodules changes, however, it does not update the submodules.
> git pull

To finalize the update, you need to run git submodule update
> git submodule update --init --recursive

## Features
- Submit a message for filtering
- Fetch rejected messages by any of the validators 
- Fetch queued messages (messages that need manual image approval)

## Architecture and design decisions
### API
- Home page used for testing connection to the service
> GET "/" 
> > Welcome to the filter service!!

- Filter message route used for submitting a JSON encoded message in the Body. The message should have 2 keys:
  - id (string)
  - body (string) - in the Markdown format 
> POST "/filter_message"
>> {  
     "id" : "32s333422",  
     "body": "#Heading 1<br> paragraph<br> [link](path/to/image)"  
  }

- Rejected message list has the format given below with the keys id(string), body(string), state(string), reason(string).
> GET "/rejected_messages"
>> {  
"success": true,  
"response": [  
{    
"id": "32s33322",  
"body": "#Heading 1\ncat paragraph ",  
"State": "Rejected",  
"Reason": "LanguageValidationFailed"  
},  
{  
"id": "32s33342214d",  
"body": "#Heading 1\ncat paragraph ",  
"State": "Rejected",  
"Reason": "LanguageValidationFailed"  
}]}

- Queued message list has the format given below with the keys id(string), body(string), state(string), reason(string).
> GET "/queued_messages"
>> {  
"success": true,
"response": [  
{  
"id": "32s333422",  
"body": "#Heading 1\n paragraph\n ![alt image](path/to/image)",  
"State": "Queued",  
"Reason": "ManualValidationNeeded"  
}]}


### Structure
The brain of the service is called `processor`. This component is used to decouple interfaces from concrete implementations.  
  
The processor contains:
- A list of `validator interfaces` (validators.Base) having a Validate method. This means that the processor doesn't need to know what concrete types will be added in the list. His only job is to parse the list, call `Validate` and get the result.
- A repository (that is also an interface) which the processor uses to store messages based on the result of the validator.  

Having this approach allowed the service to easily create concrete types of validators, each implementing validators.Base interface. Furthermore, the service can remove or add validators without any modifications in the processor. Regarding the repository, the idea is similar, meaning that the service can easily switch from a database repository to another type of repository like file based repository as long as it implements the interface.
> []validators.Base{validators.NewLinkValidator(), validators.NewImageValidator(), validators.NewTextValidator(), validators.NewLanguageValidator(repo)}

The `validators` are:
- **ImageValidator**
  - Responsible for checking if the message has an image in the following format \!\[alt](path)
  - If the message contains an image, the validator will set `State: Queued` and `Reason: ManualValidationMeaning`
  - After that, the processor is responsible for storing the message in the `QueuedMessages table`.
  - If the message contains an `image with tags` (meaning that it was already manually validated) it verifies the tags. If there is at least one tag with Rejected, the message has `State: Rejected`, otherwise `State: Approved`
  - Tags have the comment syntax from Markdown: \<!--state: Accepted--> \<!--state: Rejected-->
  - If the message does not contain any images, it has `State: Approved`
- **LinkValidator**
  - Checks if the message contains any links
  - Links have the following Markdown syntax \[name](link)
  - If the link is external, the validator will set `State: Rejected`and `Reason: LinkValidationFailed`, otherwise the message will have `State: Approved`
  - If message does not contain any links it will have `State: Approved`
- **TextValidator**
  - Checks if the message starts with a level 1 heading and has at least a paragraph.
  - If the message contains the expected elements it will set `State: Approved`, otherwise
  - Line breaks supported by the validator are \<br> and \n.
- **LanguageValidator**
  - This validator is a bit more complex than the other because it needs to check if the message contains any `banned words`. 
  - The `banned words list` comes from a different service called language-service.
  - This service provides subscription based updates on a given URL.
  - To use the language-service, the filter-service must subscribe by making a `POST to /subscribe` route provided by the language-service API. The POST should contain a `Token` and a `URL`
  - When the language-service receives an updated list, it is responsible for notifying all the subscribers at the given `URL` with the new `banned words list` and the `Token`. The `Token` is used to make sure that we receive the list from a trusted source. The `URL` and the `Token` used by filter-service to subscribe are found in the config.json file.
  - In other words, the services communicate using the Observer Pattern adapted for web applications.

## Further improvements
- The reason could be more specific like MissingHeader, MissingParagraph.
- Improve regex to match all the possible cases, including HTML syntax.
- Use a logging framework
- Control log level
- Improve testing coverage
