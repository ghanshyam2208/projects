import { ICustomerRepository } from 'data/repository/i-customer-repository';
import { injectable, inject } from 'inversify';
import 'reflect-metadata';

@injectable()
export class CustomerService {
  constructor(
    @inject('ICustomerRepository')
    private customerRepository: ICustomerRepository,
  ) {}
  getAll() {
    return this.customerRepository.getAll();
  }
  get() {
    return this.customerRepository.get();
  }
}
