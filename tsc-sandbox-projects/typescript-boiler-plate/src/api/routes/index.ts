import { Application, Request, Response } from 'express';
import v1 from './v1';

export class Routes {
  public routes(app: Application): void {
    app.route('/status').get((req: Request, res: Response) => {
      res.status(200).send('Healthy!!!');
    });
    app.use('/v1', v1);
  }
}
