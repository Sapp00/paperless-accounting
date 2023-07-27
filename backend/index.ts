import express, { Express, Request, Response } from 'express';
import { APP_PORT } from './settings';
import { allExpenses } from './routes/expenses';

const app: Express = express();

app.get('/', allExpenses);

app.listen(APP_PORT, () => {
  console.log(`[server]: Server is running at http://localhost:${APP_PORT}`);
});