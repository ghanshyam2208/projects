import { CustomerService } from '../api/services/customer.service';
import { CustomerRepository } from '../data/repository/customer.repository';
import { ICustomerRepository } from '../data/repository/i-customer-repository';
import { Container } from 'inversify';
export class DiContainer {
  public diContainer: Container;

  constructor() {
    this.diContainer = new Container();
    this.diContainer
      .bind<CustomerService>(CustomerService)
      .toSelf()
      .inSingletonScope();
    this.configure();
  }
  public configure() {
    this.diContainer
      .bind<ICustomerRepository>('ICustomerRepository')
      .to(CustomerRepository)
      .inSingletonScope();
  }
}

export default new DiContainer().diContainer;
