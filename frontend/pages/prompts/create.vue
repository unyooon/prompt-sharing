<template>
  <div class="create-prompt">
    <div class="create-prompt__title">
      <h1>Create Prompt</h1>
    </div>
    <form @submit.prevent="handleSubmit">
      <div class="create-prompt__form-group">
        <label class="create-prompt__label" for="description">Title</label>
        <FormsInput
          type="text"
          id="description"
          v-model="form.description"
          placeholder="Title"
          required
        />
      </div>
      <div class="create-prompt__form-group">
        <label class="create-prompt__label" for="llmId">LLM</label>
        <FormsSelect
          :options="llms"
          :selected-option-id="form.llmId"
          @select-option="selectLlms"
        />
      </div>
      <div class="create-prompt__form-group">
        <label class="create-prompt__label" for="text">Prompt</label>
        <FormsMultiTextInput
          type="textarea"
          id="text"
          v-model="form.text"
          placeholder="Generate a 500-word essay discussing the impact of artificial intelligence on modern education. Include examples of AI applications in both K-12 and higher education settings."
          required
        />
      </div>
      <!-- <div class="create-prompt__form-group">
        <label class="create-prompt__label" for="parameters">パラメータ</label>
        <div
          v-for="(parameter, index) in form.parameters"
          :key="index"
          class="parameter-group"
        >
          <FormsInput
            type="number"
            v-model="parameter.parameterId"
            placeholder="パラメータID"
            required
          />
          <FormsInput
            type="number"
            v-model="parameter.value"
            placeholder="パラメータ値"
            required
          />
          <button type="button" @click="removeParameter(index)">削除</button>
        </div>
        <button type="button" @click="addParameter">パラメータ追加</button>
      </div> -->
      <div class="create-prompt__submit">
        <ButtonsFillButton text="Create" />
      </div>
    </form>
  </div>
</template>

<script lang="ts" setup>
import type { CreatePromptRequest, Parameter } from "~/types/api";

const router = useRouter();
const { prompts } = useApi();

const form = ref<CreatePromptRequest>({
  description: "",
  llmId: 2,
  text: "",
  parameters: [],
});

// TODO: llmはAPIから取得
const llms = ref([
  { id: 1, name: "LLM 1" },
  { id: 2, name: "LLM 2" },
  { id: 3, name: "LLM 3" },
]);

const addParameter = () => {
  form.value.parameters.push({ parameterId: 0, value: 0 });
};

const removeParameter = (index: number) => {
  form.value.parameters.splice(index, 1);
};

const selectLlms = (llmId: number) => {
  form.value.llmId = llmId;
};

const handleSubmit = async () => {
  try {
    await prompts.post(form.value);
    router.push("/prompts");
  } catch (error) {
    console.error("Error creating prompt:", error);
  }
};
</script>

<style lang="scss" scoped>
h1 {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1.75rem;
}

.create-prompt {
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  border: 1px solid #ccc;
  border-radius: 8px;

  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas,
    Liberation Mono, Courier New, monospace;

  &__title {
    width: 100%;
    display: flex;
    justify-content: center;
  }

  &__label {
    display: block;
    margin-bottom: 0.5rem;
  }

  &__form-group {
    margin-bottom: 2rem;
  }

  &__submit {
    display: flex;
    justify-content: center;
  }
}

.parameter-group {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}
</style>
