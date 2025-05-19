// controllers/account_controller.js
class UserController {
  constructor({ accountClient, userRepository }) {
    this.accountClient = accountClient;
    this.userRepository = userRepository;
  }

  async getAccounts(req, res) {
    try {
      const accounts = await this.userRepository.findAll();
      res.json(accounts);
    } catch (err) {
      res.status(500).json({ error: err.message });
    }
  }
}
function NewUserController(accountClient, userRepository) {
  return new UserController({ accountClient, userRepository });
}
module.exports = { NewUserController };
