gohttpapidoc
=========================

Create a markdown file with the documentation of your api.

You have to comment your handlers like this :

    /*
    HTTP API
        - Action : Description of the method
        - Method : POST
        - Url : /endpoint
        - Params : ?param1=value
        - Return success : 200 { type : "ok" }
        - Return error :   400 { type : "error", message : "The error message" }
    */
