import express, { Application, Request, Response, Router } from 'express';
import { CustomerRoute } from './customer.route';

export class AppRoutes {
  constructor() {}
  public setRoutes(app: Application) {
    app.get('/status', (_req: Request, res: Response) => {
      return res.status(200).send();
    });
    // initialize customer router
    const customerRouter = new CustomerRoute();
    app.use('/customers', customerRouter.setCustomerRouter());
  }
}
