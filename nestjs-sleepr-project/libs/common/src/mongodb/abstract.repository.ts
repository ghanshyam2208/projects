import { FilterQuery, Model, Types, UpdateQuery } from 'mongoose';
import { AbstractDocument } from './abstract.schema';
import { Logger, NotFoundException } from '@nestjs/common';

/**
 * AbstractRepository
 *
 * Abstract class that provides base functionality
 * for all MongoDB repositories
 *
 * @template TDocument - document type
 */
export abstract class AbstractRepository<TDocument extends AbstractDocument> {
  protected abstract readonly logger: Logger;
  constructor(protected readonly model: Model<TDocument>) {}

  /**
   * Creates a new document.
   *
   * @param {Omit<TDocument, '_id'>} document - the document to be created
   * @return {Promise<TDocument>} a promise that resolves to the created document
   */
  async create(document: Omit<TDocument, '_id'>): Promise<TDocument> {
    const createdDocument = new this.model({
      ...document,
      _id: new Types.ObjectId(),
    });

    return (await createdDocument.save()).toJSON() as unknown as TDocument;
  }

  /**
   * async findOne function to find a document based on the filter query.
   *
   * @param {FilterQuery<TDocument>} filterQuery - the filter query to find the document
   * @return {Promise<TDocument>} the found document
   */
  async findOne(filterQuery: FilterQuery<TDocument>): Promise<TDocument> {
    const document = await this.model
      .findOne(filterQuery)
      .lean<TDocument>(true);

    if (!document) {
      this.logger.warn('document not found with filter query ', filterQuery);
      throw new NotFoundException('document not found with filter query ');
    }

    return document;
  }

  /**
   * Async function to find and update a document.
   *
   * @param {FilterQuery<TDocument>} filterQuery - The filter query to find the document to update
   * @param {UpdateQuery<TDocument>} updateQuery - The update query to apply to the document
   * @return {Promise<TDocument>} The updated document
   */
  async findOneAndUpdate(
    filterQuery: FilterQuery<TDocument>,
    updateQuery: UpdateQuery<TDocument>,
  ): Promise<TDocument> {
    const document = await this.model
      .findOneAndUpdate(filterQuery, updateQuery, {
        new: true, // mongoose returns a new/updated document in response
      })
      .lean<TDocument>(true);

    if (!document) {
      this.logger.warn('document not found with filter query ', filterQuery);
      throw new NotFoundException('document not found with filter query ');
    }

    return document;
  }

  /**
   * Find multiple documents based on the provided filter query.
   *
   * @param {FilterQuery<TDocument>} filterQuery - the query used to filter the documents
   * @return {Promise<TDocument[]>} a promise that resolves to an array of documents
   */
  async findMany(filterQuery: FilterQuery<TDocument>): Promise<TDocument[]> {
    return this.model.find(filterQuery).lean<TDocument[]>(true);
  }

  /**
   * Finds a single document and deletes it based on the provided filter query.
   *
   * @param {FilterQuery<TDocument>} filterQuery - the filter query to find the document to delete
   * @return {Promise<TDocument>} a promise that resolves to the deleted document
   */
  async findOneAndDelete(
    filterQuery: FilterQuery<TDocument>,
  ): Promise<TDocument> {
    return this, this.model.findOneAndDelete(filterQuery).lean<TDocument>(true);
  }
}
