export function useClipboard(timeoutDelay = 2000) {
  let copiedId = $state<string | null>(null);

  const copyText = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      copiedId = text;
      
      setTimeout(() => {
        if (copiedId === text) copiedId = null;
      }, timeoutDelay);
    } catch (err) {
      console.error("Failed to copy: ", err);
    }
  };

  return {
    get copiedId() {
      return copiedId;
    },
    copyText
  };
}