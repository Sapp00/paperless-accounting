import { error } from "console";
import { PAPERLESS_EXPENSE_TAG, PAPERLESS_URL, PAPERLESS_AUTH_TOKEN, PAPERLESS_UNSAFE_SSL } from "../settings"

type PaperlessDocument = {
    id: number;
    correspondent: number; 
    title:string;
    content:string;
    tags:string[];
    created_date:Date;

 /*   constructor(id: number, correspondent: number, title:string, content:string, tags:[string], created_date:Date){
        this.id = id;
        this.correspondent = correspondent;
        this.title = title;
        this.content = content;
        this.tags = tags;
        this.created_date = created_date;
        
    }*/
}
type PaperlessDocumentResponse = {
    count: number;
    next: string;
    previous: string;
    all: number[];
    results:PaperlessDocument[]
}

async function paperlessDocumentQueryExecute(query: string, options: object, pageNumber: number): Promise<PaperlessDocument[]>{
    let uri = (new URL(`/api/documents/?page=${pageNumber}&query=${query}`, PAPERLESS_URL));

   // console.log(uri);
    let response = await fetch(uri, options );
    if(!response?.ok){
        throw error(`Fetch error: ${response?.status}`);
    }
    let data: PaperlessDocumentResponse = await response.json();

    // preprocess data
    let res = data.results;
    res.forEach( (r) => r.created_date = new Date(r.created_date));

    console.log(`length: ${res.length}`);
    if (data.next != null){
        pageNumber++;
        let response = await paperlessDocumentQueryExecute(query, options, pageNumber) ;

        res = res.concat ( response );
    }

    
    return res;
}

async function paperlessDocumentQuery(query: string): Promise<PaperlessDocument[]>{
        try{
            const options = {
                method: "GET",
                headers: {
                Accept: "application/json",
                "Content-Type": "application/json;charset=UTF-8",
                "Authorization": "Token "+PAPERLESS_AUTH_TOKEN
                },
            };
            var all_expenses: PaperlessDocument[] = [];

            let uri = (new URL("/api/documents/?query="+query, PAPERLESS_URL));
            process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

            let resp = await paperlessDocumentQueryExecute(query, options, 1);

            return resp;
        }
        catch(e: any){
            console.error(e);
            return [];
        }
}

export {PaperlessDocument, paperlessDocumentQuery};