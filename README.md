# Pwitter
A web application to be used as a benchmarking tool for IaaS systems.  
It exposes RESTfull endpoints to perform both CPU-intensive and memory-intensive tasks.

### What Does it Do?
__Pwitter__ is similar to Twitter, with the difference that _pweets_ (tweets) are not indexed by their hashtags, but by their _sentiment_ (_this could change in future_).

Thus, __Pwitter__, stores _pweets_ together with their [sentiment analisys](https://textblob.readthedocs.org/en/dev/index.html).  
_Pweets_ are requested by their sentiment.

## The Endpoints
