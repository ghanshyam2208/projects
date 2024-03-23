import { Application, Request, Response } from 'express';
import customerRouter from './customer.route';

export class AppRoutes {
  public setRoutes(app: Application) {
    app.get('/status', (_req: Request, res: Response) => {
      return res.status(200).send();
    });
    app.use('/customer', customerRouter);
  }
}
