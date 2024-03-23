import express, { Request, Response, Router } from 'express';

export class CustomerRoute {
  public router: Router;
  public routerPath = 'customers';
  constructor() {
    this.router = express.Router();
  }

  public setCustomerRouter() {
    this.router.get(``, (_req: Request, res: Response) => {
      return res.status(200).json({ h1: 'h1' });
    });

    return this.router;
  }
}
