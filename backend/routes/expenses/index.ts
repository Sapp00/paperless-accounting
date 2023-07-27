import { PaperlessDocument, paperlessDocumentQuery } from "../../core/paperlessDocument";
import { PAPERLESS_EXPENSE_TAG } from "../../settings";
import {Request, Response } from 'express';


async function allExpenses(req: Request, res: Response){
    type ChartEntry = {
        date: number,
        category: string,
        value: number
    };

    // TODO: here is potential for performance optimization
    try{
        const all_expenses: PaperlessDocument[] = await paperlessDocumentQuery("tag:"+PAPERLESS_EXPENSE_TAG);

        var expense_chart_dict = new Map<number, ChartEntry>();
        var expense_chart_paid_dict = new Map<number, ChartEntry>();
        var expense_sum = 0;
        var paid_sum = 0;

        // create entries
        all_expenses.forEach( (e) => {
            // TODO: change created date! needs to be based on paid_date which is retrieved from the database
            let e_price = e.title.charCodeAt(0)*20;
            let e_paid = (e_price % 3 != 0) ? e_price : 0;
            let e_paid_date = e.created_date.getTime();
            let e_date = e.created_date.getTime();

            paid_sum += e_paid;
            expense_sum += e_price;
            if (e_date in expense_chart_dict){
                expense_chart_dict.get(e_date)!.value = expense_sum;
            } else {
                expense_chart_dict.set(e_date, {'date': e_date, 'category': 'expense', 'value': expense_sum} );
            }

            if (e_paid_date in expense_chart_paid_dict) {
                expense_chart_paid_dict.get(e_paid_date)!.value = paid_sum;
            } else {
                expense_chart_paid_dict.set( e_paid_date, {'date': e_date, 'category': 'payment', 'value': paid_sum} );
            }
        });

        // merge expenses+income
        let expense_chart_data: ChartEntry[] = [];
        expense_chart_dict.forEach( (value) => {
            expense_chart_data.push(value);
        });
        expense_chart_paid_dict.forEach( (value) => {
            expense_chart_data.push(value);
        });

        res.status(200);
        res.send(expense_chart_data);
    }
    catch(e: any){
        console.error(e);
        res.status(400);
        res.send('Error');
    }
}

export {allExpenses};