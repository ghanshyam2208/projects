import express, { Response, Request } from 'express';
import { UserController } from '../../../controller';

const router = express.Router();

router.get('/profile', UserController.getUserProfileData);

export default router;
