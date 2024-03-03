import { Prop, Schema } from '@nestjs/mongoose';
import { SchemaTypes, Types } from 'mongoose';

@Schema()
/**
 * AbstractDocument
 *
 * Base schema for all MongoDB documents
 * @see https://docs.nestjs.com/techniques/mongodb
 */
export class AbstractDocument {
  /**
   * ID of the document in MongoDB
   */
  @Prop({ type: SchemaTypes.ObjectId })
  _id: Types.ObjectId;
}
