import express, { Router, Request, Response } from 'express';

class CustomerRoutes {
  public router: Router;
  constructor() {
    this.router = express.Router();
    this.setRoutes();
  }

  private setRoutes() {
    this.router.get(``, (_req: Request, res: Response) => {
      return res.status(200).json({ h1: 'h1' });
    });
  }
}

export default new CustomerRoutes().router;
