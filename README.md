# GoDeepl - A Deepl API wrapper

This is a Go wrapper around the *(mostly undocumented)* JSON RPC API of the online translation service [Deepl](https://www.deepl.com/translator).

## How to use it

### Authentication

Create a client specifying the preferred endpoint to be used. You can chose between the public API *(which is strongly ratelimited)* and the pro API which needs authentication.

#### Public API

```go
// Defaultly, when no options are passed, 
// EndpointPublic is used as API endpoint.
c := godeepl.New()
```

#### Pro API

```go
c := godeepl.New(godeepl.ClientOptions{
    // Select the Pro API endpoint
	Endpoint: godeepl.EndpointPro,

    // Login Credentials.
    // When left empty, you need to pass a pre-obtained
    // SessionID.
	Email:    os.Getenv("EMAIL"),
	Password: os.Getenv("PASSWORD"),

	// When the SessionID is empty, Email and Password is used
	// for login and the obtained sessionID is stored in the Client.
	SessionID: os.Getenv("SESSION"),
})
```

### Translate Text

Now, you can use the client to translate text.

You can also pass additional options to the translate endpoint like the desired formality
or the ammount of beams (translation alternatives).

```go
res, err := c.Translate(
	godeepl.LangAuto, godeepl.LangEnglish,
	"Jo, was geht ab du alter Sack! Dauert das noch lange?",
	godeepl.TranslationOptions{
		Formality:        godeepl.FormalityFormal,
		PreferedNumBeams: 3,
	})
```

The returned `TranslationResult` is an array of translations for each sentence, which again
contains a list of beams. There are some utility functions on this object to simplify
obtaining the translation results.

For more details, please take a look into the provided [examples](examples).

---

© 2021 Ringo Hoffmann (zekro Development).  
Covered by the MIT License.