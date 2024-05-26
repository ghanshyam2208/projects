import { Test } from '@nestjs/testing';
import { UserModule } from 'src/user.module';

describe('UsersController', () => {
  beforeAll(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [UserModule],
    }).compile();

    const app = moduleRef.createNestApplication();
    await app.init();
  });
});
