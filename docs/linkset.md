# Link Sets

## About

Need to review the linkset aspect and how to read these and potentially feed them
into a cralwer approach. 

```json
{
    "linkset": [
    {
        "rel": "self",
        "href": "https://example.org/links/resource1"
    },
    {
        "rel": "alternate",
        "href": "https://example.org/links/resource1.html"
    },
    {
        "rel": "related",
        "href": "https://example.org/links/resource1.json"
    }
    ]
}
```

```text
<https://example.org/links/resource1>; rel="self"
<https://example.org/links/resource1.html>; rel="alternate"
<https://example.org/links/resource1.json>; rel="related"
```