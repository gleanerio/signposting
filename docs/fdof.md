# FDO Filesystem concept

## About

Likely should use the FDOF-IR


```turtle
<fdofirIdentifier> a fdof-o:fodfIR
fdof-o:isMetadataOf <fdoIdentifier>.
<fdoIdentifier> fdof-o:hasType <FDOType>
fdof-o:hasMetadata <metadataRecordIdentifier>
fdof-o:hasObjectLocation <ObjectLocation>
```

See also the section 3.1.3 FDOF-P using HTTP Signposting

```
HTTP/1.1 200
Server: nginx/1.21.6
Content-Type: text/turtle
Connection: keep-alive
Link: <https://fdof.org/object1/fdof-ir> ; rel="fdof-ir",
<https://fdof.org/object1/metadatarecord1> ; rel="fdof-metadata" ,
<https://fdof.org/object1/metadatarecord2> ; rel="fdof-metadata" ,
<https://fdof.org/types/dataset> ; rel="fdof-type"
Accept: text/html, text/turtle, application/ld+json, application/rdf+xml
Expires: 0
```


## References

* [FDOF](https://fairdigitalobjectframework.org/)
