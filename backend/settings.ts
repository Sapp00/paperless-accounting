import dotenv from 'dotenv';

dotenv.config();

const APP_PORT = process.env.APP_PORT;
const PAPERLESS_EXPENSE_TAG: string = process.env.PAPERLESS_EXPENSE_TAG!;
const PAPERLESS_INCOME_TAG:string = process.env.PAPERLESS_INCOME_TAG!;
const PAPERLESS_URL:string = process.env.PAPERLESS_URL!;
const PAPERLESS_AUTH_TOKEN:string = process.env.AUTH_TOKEN!;
const PAPERLESS_UNSAFE_SSL:boolean = (process.env.PAPERLESS_URL == 'true');

export {APP_PORT, PAPERLESS_EXPENSE_TAG, PAPERLESS_INCOME_TAG, PAPERLESS_URL, PAPERLESS_AUTH_TOKEN, PAPERLESS_UNSAFE_SSL};