import type { ReadLlmMasterResponse } from "~/types/api";

export const llmMastersApi = () => {
  // fetchLLMs関数をエクスポート
  const get = async (): Promise<Array<ReadLlmMasterResponse> | null> => {
    const url = "/llm-masters"; // APIのURLを構築
    const result = await useCustomFetch<Array<ReadLlmMasterResponse>>(url); // カスタムフェッチを実行

    return result; // データを返す
  };

  return { get };
};
