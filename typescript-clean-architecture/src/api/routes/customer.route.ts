import express, { Router, Request, Response } from 'express';

class CustomerRoutes {
  public customerRouter: Router;
  constructor() {
    this.customerRouter = express.Router();
    this.setCustomerRoutes();
  }

  private setCustomerRoutes() {
    this.customerRouter.get(``, (_req: Request, res: Response) => {
      return res.status(200).json({ h1: 'h1' });
    });
  }
}

export default new CustomerRoutes().customerRouter;
