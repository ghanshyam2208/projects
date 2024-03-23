import express, { Router, Request, Response } from 'express';
import diContainer from '../../Di/di-container';
import { CustomerService } from '../services/customer.service';

class CustomerRoutes {
  public router: Router;
  public customerService: CustomerService;
  constructor() {
    this.router = express.Router();
    this.customerService = diContainer.get<CustomerService>(CustomerService);
    this.setRoutes();
  }

  private setRoutes() {
    this.router.get(``, (_req: Request, res: Response) => {
      return res.status(200).json({ h1: this.customerService.get() });
    });
  }
}

export default new CustomerRoutes().router;
