import { UserServices } from '../services';
import { Request, Response } from 'express';

export class UserController {
  public static getUserProfileData(req: Request, res: Response) {
    try {
      const data = UserServices.getUserProfileData();
      res.send(data).status(200);
    } catch (error) {
      console.log(error);
      throw error;
    }
  }
}
