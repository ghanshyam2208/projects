import { AbstractRepository } from '@app/common/database/abstract.repository';
import { UsersDocument } from './users.model';
import { Model } from 'mongoose';
import { Logger } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';

export class UsersRepository extends AbstractRepository<UsersDocument> {
  protected readonly logger = new Logger(UsersRepository.name);

  constructor(
    @InjectModel(UsersDocument.name)
    usersModel: Model<UsersDocument>,
  ) {
    super(usersModel);
  }
}
