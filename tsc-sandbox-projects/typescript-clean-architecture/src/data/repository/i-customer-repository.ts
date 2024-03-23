export interface ICustomerRepository {
  get(): string;
  getAll(): string[];
}

const TYPES = {
  ICustomerRepository: Symbol.for('ICustomerRepository'),
};

export { TYPES };
