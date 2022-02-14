## QuickShare

Lightweight program to share files as well as generating links to send to people so that they can upload files to your
server.

### Service Features

#### Upload
* URL Generation
  * Create URL which is shared with person who you want to upload file to your server
    * Generate token
    * Store token
      * redis
      * db
      * flat file
  * URL allows person to upload one file
    * Define allowable MIME types
      * Images, video files, PDFs, docx, etc
  * Token is invalidated after successful upload.

#### Download
* Generate link to a specific file
  * Token Expiration
    * Time elapsed since link creation
    * Number of downloads


### Work Needed

* DB
  * Schema
  * Flavor

* Dockerizing it
 * This should probably be done prior to anything else.

* Config
  * Read in from file (json, yaml, toml, etc)
  * Store in running program
  * Values
    * Port
    * Domain (Include subdomain, if applicable) Maybe leave to rev proxy to deal with this.
    * Store/DB connection values
      * Domain
      * Port
      * Credentials
* File Object
  * JSON-serializable struct
  * Fields
    * id string
    * Valid bool
    * Expires int unix epoch timestamp
    * 
* CLI
  * Manually create links
    * `qs-cli create-upload`
      * Creates token string that isn't in store 
      * Stores token string
      * Outputs `qs.example.tld/upload?token=1234567890` to terminal
    * `qs-cli share $FILE`
      * Copies files to share directory
      * Create token 
  * Does not attempt to listen on http servers, just talks to store to create tokens
* Web Interface
  * File upload page
    * File picker, upload button
    * Upload progress bar
    * Success / fail dialog
  * Invalid or expired token page.
  * Generate upload / download links
* Webserver
  * Listen on port
  * Extract token from url