#!/usr/bin/env node

import { Observable, empty, from, of } from "rxjs";
import { flatMap } from "rxjs/operators";
import axios from "axios";
import commander from "commander";

import { validateRequest } from "./model";

let obs = Observable.create( (observer) => {
    process.stdin.
        on('data', data => observer.next(data) ).
        on('end', x => observer.complete() ).
        on('error', err => observer.error(err) );
});

var buff = "";
const toLines = flatMap( data => {
    buff += data;
    let lines = buff.split(/[\r\n|\n]/);
    buff = ""+lines.pop();
    if(lines.length > 0) {
        return from(lines);
    } else {
        return empty();
    }
} );

const toJson = flatMap( (line:string) => {
    try{
        return of(JSON.parse(line));
    }catch(error){
        process.stderr.write("ERR: "+error+"\n");
        return empty();
    }
});

const validate = flatMap( request => {
    let err = validateRequest(request);
    if(err!=null){
        process.stderr.write("ERR: "+err+"\n");
        return empty();
    }
    return of(request);
});

const doIt = flatMap( (reqres:any) => {
    return Observable.create( observer => {
        try{
            let request = reqres["request"];
            axios({
                url: request["url"],
                method: request["method"]
            }).then(res => {
                if(reqres["response"] != null){
                    if(reqres["_meta"] == null){
                        reqres["_meta"] = {};
                    }
                    reqres["_meta"]["original"] = reqres["response"]
                }
                let buf = Buffer.from(res.data);
                reqres["response"] = {
                    "status": res.status,
                    "reason": res.statusText,
                    "headers": res.headers,
                    "body": buf.toString("base64")
                };
                observer.next(reqres);
                observer.complete();
            }).catch(err => {
                console.log(err);
                observer.complete();
            });
        }catch(error){
            console.log(error);
            observer.complete();
        }
    } );
});

commander
    .version("0.0.1")
    .description("need request, do request and fill responses")
    .parse(process.argv)
obs.pipe(toLines, toJson, validate, doIt)
    .subscribe((x) => {
        process.stdout.write(JSON.stringify(x));
        process.stdout.write("\n");
    });

