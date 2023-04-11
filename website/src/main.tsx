// Dependencies (Root) and application structure (App) will be loaded in parallel
Promise.all([import('@/Root'), import('@/App')]).then(([{ default: render }, { default: App }]) => {
  render(App);
});
export {};
