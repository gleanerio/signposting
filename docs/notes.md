# Notes

## About

The following are just some personal observations and thoughts.

### Question 1

Signposting can be done via HTTP and HTML approaches.  The later
is obviously easier for a group to implement since it doesn't rely
on the need for a server that can modify the HTTP header based on a
resource to be served.

Modifying the HTML header with either signposting or JSON-LD
script tag approach seems an equal lift.  However, the JSON-LD is a full
data graph while signposting seems a bit harder to implement arbitrary data.

Things like variable measured or geospatial data seem easier in a JSON-LD
data graph vs signposting link authoring.  Is this a case of
Structured Data on the Web (SDOTW) being
more for knowledge graph representation vs signposting be ing more for
PID graph represenation?

It would be interesting to see how Signposting might directly
reference a JSON-LD document allowing for fast indexing and PID graph
generation which is then augmented with a second pass for the more
verbose KG graph focused JSON-LD documents.

Signposting obviously, it seems, does this via things like the example
at [32-http-describedby-profile-conneg
](https://s11.no/2022/a2a-fair-metrics/32-http-describedby-profile-conneg/).

Can we view signposting as both an alternative to SDOTW when a more
PID graph generating approach is desired?  Additionally then, signposting
could be seen as an evolved SDOTW approache that provides alternative
linking, profiling, etc for the core SDOTW approach of JSON-LD in
script tags with sitemaps and robots.txt.  

![image](./images/concept.svg)

In the above the resources are still defined by things
like [sitemap.xml](https://sitemaps.org/) or
[robots.txt](https://www.rfc-editor.org/rfc/rfc9309.html). The nature of
Signposting does seem allow for an easier implmentation
of a classic crawler if desired.  This aspect of controlling how
indexing is done does seem to be getting some [fresh attention](https://blog.google/technology/ai/ai-web-publisher-controls-sign-up/)


### Question 2

With the emergance of JSON-LD for the SDOTW approach, a huge ecosystem of
tooling is leveraged since JSON-LD is just JSON. This does effectively hide
some of the RDF elements in more classic LOD based approaches.

There is ample toolinmg for thisn like validation (SHACL, JSON Schema, Shex),
schema (JSON Schema), query tools, authoring libraries and much more are all
available.  The resulting JSON(-LD) is then loaded into a single script tag
rather than multiple link elements. 

For publishers, signposting involves the creation of a method to map
from their local data store to HTML link headers.  Which does not seem as
mature as those for serializing to JSON strcutures seems.   This could also
just be me not aware of them


### Question 3

The Signposting approach leveraging HTTP headers does have the potential for
very large gains in performance and efficency for indexing large records.
It definately seems there are different used cases that can lead to
different architectural selections.

### Question 4

It is interesting to look at some of the various approaches that groups
are using for discovery.  The above points revolve around SDOTW and Signposting
but other groups are using things like and STAC catalogs mapped to SDOTW,
MQTT pubsub approaches.

Note also that some of the Signposting authors view it as a way to represent
the topology of research objects.  The manner in which is enables or improves
the crawling of this ecosystem is unknown to the author.

### Question 5

What is the subject of the Singposting statements?  Is the URL of the source
considered something like the  _subject IRI_ of the link statements?
I ask only since trying to map these to a graph I need that concept.  Though
I have not fully thought this trhough to an implementation so any guidance is
appreciated. 

Or is this is the case where a FDOF-IR is required to convey this information
and if the FDOF-IR is not present then any conversion of links to triples is 
problematic or at least ad hoc.  


```JSON
{
  "@context": {
    "@vocab": "https://example.org/id/fdof-o/"
  },
  "@graph": [
    {
      "@id": "http://njh.me/fdoIdentifier",
      "hasMetadata": {
        "@id": "http://njh.me/metadataRecordIdentifier"
      },
      "hasObjectLocation": {
        "@id": "http://njh.me/ObjectLocation"
      },
      "hasType": {
        "@id": "http://njh.me/FDOType"
      }
    },
    {
      "@id": "http://njh.me/fdofirIdentifier",
      "@type": "fodfIR",
      "isMetadataOf": {
        "@id": "http://njh.me/fdoIdentifier"
      }
    }
  ]
}
```

The above IR would then be referenced by rel  _fdof-ir_ I assume?
As seen below

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

I am not yet sure I fully understand the relationship between the IR RDF
and RDF represented by 

```
<https://fdof.org/object1/metadatarecord1> ; rel="fdof-metadata" ,
```

which might be typed as _application/ld+json_.

I assume the IR should be limitd in scope only to the IR and is there a 
suggsted namespace for it?   Or can one simple leverage DCAT or schema.org for 
example?   


### References

* [CDIF Discovery](https://github.com/Cross-Domain-Interoperability-Framework/Discovery)
* [FAIR Digital Object Framework Documentation](https://fairdigitalobjectframework.org/)
* [FAIR Digital Objects and FAIR Signposting](https://zenodo.org/record/7977333)
* [JSON-LD](https://json-ld.org/) and [Data on the Web Best Practices](https://www.w3.org/TR/dwbp/)
* [Signposting](https://signposting.org/)
* [W3C Content Negotiation by Profile](https://www.w3.org/TR/dx-prof-conneg/)
* [STAC-Browser](https://github.com/radiantearth/stac-browser) 
* [MQTT](https://mqtt.org/)
* [GeoAPI](https://www.geoapi.org/) with [Java](https://www.geoapi.org/3.0/javadoc/org.opengis.geoapi/module-summary.html) and [Python / pygeoapi](https://pygeoapi.io/) implementations
* [Open Archives Initiative ResourceSync Framework Specification](https://www.openarchives.org/rs/toc) and [spec](https://www.openarchives.org/rs/1.1/resourcesync)
* [FDP Specs](https://fairdo.org/specifications/)
* [Content Negotiation by Profile](https://www.w3.org/TR/dx-prof-conneg/)
* [Google Next Gen robots.txt](https://blog.google/technology/ai/ai-web-publisher-controls-sign-up/)

