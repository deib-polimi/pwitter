## Get Charts
To see charts in local, run:

```
$ python -m SimpleHTTPServer 8888
```

then connect to `localhost:8888`.

If you want to see them without cloning the repository and launching a web server, then go to http://affo.github.io/pwitter/.

## `.json` Files

`makecharts.js` script expects to read data from three different `.json` files:

 - `g.json` (The data on CPU and Memory while running `pwitter stress -G XXX`);
 - `p.json` (The data on CPU and Memory while running `pwitter stress -P XXX`);
 - `gp.json` (The data on CPU and Memory while running `pwitter stress -G XXX -P XXX`);

To generate those files, we recommend you to use the output of a `cpu_mem_perc` sensor of [Sy](https://github.com/affo/sy) container monitorer.  
You can fetch data running something like (in Sy's folder):

```
$ sy-agent -d # launches Sy's daemon
$ sy add <pwitters_container_name> cpu_mem_perc
$ sy listen > output.sydata

^CKeyboardInterrupt     # Ctrl + C
```

Now you can obtain the `.json` file -- that `makecharts.js` needs -- from `output.sydata` running: `./jsonifier output.sydata <label>`.  
If you run, for instance, `./jsonifier output.sydata "Output of stress -G 800"`, you will obtain an `output.json` file filled with ready-to-use data for `makecharts.js`. One last thing to do is to run `mv output.json g.json` in order to let the script find the file.  
The label specified in input to `jsonifier` script will be the title of the chart shown.

