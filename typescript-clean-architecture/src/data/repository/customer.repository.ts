import { injectable } from 'inversify';
import { ICustomerRepository } from './i-customer-repository';

@injectable()
export class CustomerRepository implements ICustomerRepository {
  get(): string {
    return 'found';
  }
  getAll(): string[] {
    const sampleArr = ['test', 'test'];
    return sampleArr;
  }
}
