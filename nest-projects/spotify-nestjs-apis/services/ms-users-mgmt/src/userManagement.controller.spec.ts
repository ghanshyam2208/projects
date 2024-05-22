import { Test, TestingModule } from '@nestjs/testing';
import { UserManagementController } from './userManagement.controller';
import { UserManagementService } from './userManagement.service';

describe('AppController', () => {
  let userManagementController: UserManagementController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [UserManagementController],
      providers: [UserManagementService],
    }).compile();

    userManagementController = app.get<UserManagementController>(
      UserManagementController,
    );
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(userManagementController.getHello()).toBe('Hello World!');
    });
  });
});
