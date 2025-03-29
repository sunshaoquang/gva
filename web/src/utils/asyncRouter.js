const viewModules = import.meta.glob("../view/**/*.vue");
const pluginModules = import.meta.glob("../plugin/**/*.vue");

export const asyncRouterHandle = (asyncRouter) => {
  asyncRouter.forEach((item) => {
<<<<<<< HEAD
    if (item.component && typeof item.component === "string") {
      if (item.component.split("/")[0] === "view") {
        item.component = dynamicImport(viewModules, item.component);
      } else if (item.component.split("/")[0] === "plugin") {
        item.component = dynamicImport(pluginModules, item.component);
=======
    if (item.component && typeof item.component === 'string') {
      item.meta.path = '/src/' + item.component
      if (item.component.split('/')[0] === 'view') {
        item.component = dynamicImport(viewModules, item.component)
      } else if (item.component.split('/')[0] === 'plugin') {
        item.component = dynamicImport(pluginModules, item.component)
>>>>>>> main
      }
    }
    if (item.children) {
      asyncRouterHandle(item.children);
    }
  });
};

function dynamicImport(dynamicViewsModules, component) {
<<<<<<< HEAD
  const keys = Object.keys(dynamicViewsModules);
=======
  const keys = Object.keys(dynamicViewsModules)
>>>>>>> main
  const matchKeys = keys.filter((key) => {
    const k = key.replace("../", "");
    return k === component;
  });
  const matchKey = matchKeys[0];

  return dynamicViewsModules[matchKey];
}
