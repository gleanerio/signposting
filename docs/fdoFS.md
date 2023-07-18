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

## Notes

Need to create FDO that define service calls into the exising NSDF data
and functions/libraries.  I can use DuckDB over a set of Parquet 
files (like the OIH graph parquet files).

### Players

- Signposting (client)
- Grow (server)
- FDO document (JSON-LD)

### Questions

- Is signposting really needed?
- The object should have a web and data representation
- This is really a case of FDO being generated to represent views into collections of data.

### Broader Impacts

- IRIS
- OIH Services

Look at previous work I did for mapping in /home/fils/src/Projects/NSDF/nsdf-catalog-graph



## References

* [FDOF](https://fairdigitalobjectframework.org/)
