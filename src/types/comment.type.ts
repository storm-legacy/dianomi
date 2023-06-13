export interface CommentReport {
  ID: number;
  ReporterID: string;
  CommentID: number;
  CreatedAt: number;
  Message: boolean;
}

export interface Comment {
  id: number;
  name: string;
  email: string;
  comment: string;
  updated_at: string;
  reports: CommentReport[]
}