import { FilterQuery, Model, Types, UpdateQuery } from 'mongoose';
import { AbstractDocument } from './abstract.schema';
import { Logger, NotFoundException } from '@nestjs/common';

export abstract class AbstractRepository<TDocument extends AbstractDocument> {
  protected abstract readonly logger: Logger;
  constructor(protected readonly model: Model<TDocument>) {}

  async create(document: Omit<TDocument, '_id'>): Promise<TDocument> {
    const createdDocument = new this.model({
      ...document,
      _id: new Types.ObjectId(),
    });
    return (await createdDocument.save()).toJSON() as unknown as TDocument;
  }

  async findOne(filterQuery: FilterQuery<TDocument>): Promise<TDocument> {
    const document = await this.model
      .findOne(filterQuery)
      .lean<TDocument>(true);

    if (!document) {
      this.logger.warn(
        `Document was not found with given filer query`,
        filterQuery,
      );
      throw new NotFoundException('Document not found');
    }
    return document;
  }

  async find(filterQuery: FilterQuery<TDocument>): Promise<TDocument[]> {
    const documents = await this.model
      .find(filterQuery)
      .lean<TDocument[]>(true);

    if (!document || !documents.length) {
      this.logger.warn(
        `Document was not found with given filer query`,
        filterQuery,
      );
      throw new NotFoundException('Document not found');
    }

    return documents;
  }

  async findOneAndUpdate(
    filterQuery: FilterQuery<TDocument>,
    updateQuery: UpdateQuery<TDocument>,
  ): Promise<TDocument> {
    const document = this.model
      .findOneAndUpdate(filterQuery, updateQuery)
      .lean<TDocument>(true);
    if (!document) {
      this.logger.warn(
        `Document was not found with given filer query`,
        filterQuery,
      );
      throw new NotFoundException('Document not found');
    }
    return document;
  }

  async findAndDelete(filterQuery: FilterQuery<TDocument>) {
    const document = await this.model
      .findOneAndDelete(filterQuery)
      .lean<TDocument>(true);

    if (!document) {
      this.logger.warn(
        `Document was not found with given filer query`,
        filterQuery,
      );
      throw new NotFoundException('Document not found');
    }
    return document;
  }
}
