/**
 * ページ情報を表すレスポンス
 */
export interface BasePageResponse {
  number: number; // ページ番号
  size: number; // ページサイズ
  totalElements: number; // 総要素数
  totalPages: number; // 総ページ数
}

/**
 * 作成されたユーザーのレスポンス
 */
export interface CreatedUserResponse {
  displayName: string; // 表示名
}

/**
 * パラメータのレスポンス
 */
export interface ParameterResponse {
  name: string; // パラメータ名
  value: number; // パラメータ値
}

/**
 * プロンプトの詳細情報を表すレスポンス
 */
export interface ReadPromptResponse {
  createdAt: string; // 作成日時
  createdUser: CreatedUserResponse; // 作成ユーザー
  description: string; // 説明
  id: number; // プロンプトID
  isPublic: boolean; // 公開フラグ
  llmName: string; // LLMの名前
  parameters: ParameterResponse[]; // パラメータのリスト
  text: string; // プロンプトのテキスト
}

/**
 * プロンプトのリストとページ情報を表すレスポンス
 */
export interface ReadPromptsResponse {
  data: ReadPromptResponse[]; // プロンプトのリスト
  page: BasePageResponse; // ページ情報
}

/**
 * プロンプト作成リクエスト
 */
export interface CreatePromptRequest {
  description: string; // 説明
  llmId: number; // LLMのID
  parameters: Parameter[]; // パラメータのリスト
  text: string; // プロンプトのテキスト
}

/**
 * パラメータ
 */
export interface Parameter {
  parameterId: number; // パラメータID
  value: number; // パラメータ値
}
