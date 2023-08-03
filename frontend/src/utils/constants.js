const name = window.FileBrowser.Name || "File Browser";
const disableExternal = window.FileBrowser.DisableExternal;
const disableUsedPercentage = window.FileBrowser.DisableUsedPercentage;
const baseURL = window.FileBrowser.BaseURL;
const logoutPath = window.FileBrowser.LogoutPath;
const staticURL = window.FileBrowser.StaticURL;
const recaptcha = window.FileBrowser.ReCaptcha;
const recaptchaKey = window.FileBrowser.ReCaptchaKey;
const signup = window.FileBrowser.Signup;
const version = window.FileBrowser.Version;
const logoURL = `${staticURL}/img/logo.svg`;
const noAuth = window.FileBrowser.NoAuth;
const authMethod = window.FileBrowser.AuthMethod;
const showLogout = window.FileBrowser.ShowLogout;
const loginPage = window.FileBrowser.LoginPage;
const theme = window.FileBrowser.Theme;
const enableThumbs = window.FileBrowser.EnableThumbs;
const resizePreview = window.FileBrowser.ResizePreview;
const enableExec = window.FileBrowser.EnableExec;
const origin = window.location.origin;

export {
  name,
  disableExternal,
  disableUsedPercentage,
  baseURL,
  logoutPath,
  logoURL,
  recaptcha,
  recaptchaKey,
  signup,
  version,
  noAuth,
  authMethod,
  showLogout,
  loginPage,
  theme,
  enableThumbs,
  resizePreview,
  enableExec,
  origin,
};
