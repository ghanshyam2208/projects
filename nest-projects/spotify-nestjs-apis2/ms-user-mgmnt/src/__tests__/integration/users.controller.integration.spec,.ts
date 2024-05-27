import { INestApplication } from '@nestjs/common';
import { Test } from '@nestjs/testing';
import { UserModule } from '../../user.module';

describe('UsersController', () => {
  let app: INestApplication;

  beforeAll(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [UserModule],
    }).compile();

    app = moduleRef.createNestApplication();
    await app.init();
  });

  afterAll(async () => {
    await app.close();
  });

  describe('Test something', () => {
    it('should return hello world', () => {
      console.log('test');
    });
  });
});
