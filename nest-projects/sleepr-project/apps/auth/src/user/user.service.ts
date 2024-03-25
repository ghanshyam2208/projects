import { CreateUserPayload } from './dto/user.payload';
import { UserRepository } from './user.repository';

export class UserService {
  constructor(private readonly userRepository: UserRepository) {}

  createUser(createUserPayload: CreateUserPayload) {
    return this.userRepository.create(createUserPayload);
  }
}
