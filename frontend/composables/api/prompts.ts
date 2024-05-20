import type { ReadPromptsResponse } from "~/types/api";

export const promptsApi = () => {
  // fetchPrompts関数をエクスポート
  const get = async (
    page: number, // ページ番号
    size: number, // ページサイズ
    llmId?: number, // LLMのID
    isPublic?: boolean // 公開フラグ
  ): Promise<ReadPromptsResponse | null> => {
    const url = "/prompts"; // APIのURLを構築
    const params = { isPublic, llmId, page, size }; // パラメータを設定
    const result = await useCustomFetch<ReadPromptsResponse>(url, {
      params,
    }); // カスタムフェッチを実行

    return result; // データを返す
  };

  return { get };
};
