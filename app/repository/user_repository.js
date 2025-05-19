class UserRepository {
  constructor({ UserModel }) {
    this.User = UserModel;
  }

  async findAll() {
    try {
      const users = await this.User.findOne();
      if (!users) {
        throw new Error('No users found');
      }
      return users;
    } catch (err) {
      throw err;
    }
  }
}

function NewUserRepository(models) {
  return new UserRepository({ UserModel: models.User });
}
module.exports = { UserRepository, NewUserRepository };
