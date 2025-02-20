import { S3Client, PutObjectCommand } from '@aws-sdk/client-s3'
import { APPLICATION_CONFIGS } from '../configs';

const STORAGE_BUCKET = APPLICATION_CONFIGS.S3_STORAGE_BUCKET;

export default class BillsFileStorage {
  static build() {
    return BillsFileStorage(
      new S3Client({ region: 'us-east-1' })
    )
  }

  constructor(client) {
    this._client = client;
  }

  async saveFromFile(filepath) {
    const command = new PutObjectCommand({
      ACL: "private",
      Bucket: STORAGE_BUCKET
    })
  }
}