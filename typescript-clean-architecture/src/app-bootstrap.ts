import { AppRoutes } from './api/routes';
import express, { Application, NextFunction, Request, Response } from 'express';
import bodyParser from 'body-parser';

export class AppBootstrap {
  public app: Application;
  public router: AppRoutes;
  constructor() {
    this.app = express();
    this.router = new AppRoutes();
    this.configureApp();
  }

  private configureApp(): void {
    this.app.use((_req: Request, res: Response, next: NextFunction) => {
      res.header('Access-Control-Allow-Origin', '*');
      res.header('Access-Control-Allow-Methods', 'GET,POST,DELETE,OPTIONS,PUT');
      res.header('Access-Control-Allow-Headers', '*');
      next();
    });
    // set up router
    this.router.setRoutes(this.app);
    // middleware allowing json requests
    this.app.use(bodyParser.json());
    this.app.use(bodyParser.urlencoded({ extended: true }));
  }
}

export default new AppBootstrap().app;
