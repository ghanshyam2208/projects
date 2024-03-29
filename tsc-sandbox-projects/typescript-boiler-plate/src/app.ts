import express, { Application, NextFunction, Request, Response } from 'express';
import bodyParser from 'body-parser';
import { Routes } from './api/routes';

class App {
  public app: Application;
  public routesObj = new Routes();
  constructor() {
    this.app = express();
    this.config();
    this.routesObj.routes(this.app);
  }

  private config(): void {
    this.app.use((req: Request, res: Response, next: NextFunction) => {
      res.header('Access-Control-Allow-Origin', '*');
      res.header('Access-Control-Allow-Methods', 'GET,POST,DELETE,OPTIONS,PUT');
      res.header('Access-Control-Allow-Headers', '*');
      next();
    });
    this.app.use(bodyParser.json());
    this.app.use(bodyParser.urlencoded({ extended: false }));
  }
}

export default new App().app;
