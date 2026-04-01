from dataclasses import dataclass
from typing import Optional, Literal
from uuid import UUID, uuid4


@dataclass
class Account:
    id: str
    owner_name: str
    balance: float

    def has_sufficient_balance(self, amount: float) -> bool:
        return self.balance >= amount

    def deposit(self, amount: float) -> None:
        self.balance += amount

    def withdraw(self, amount: float) -> bool:
        if not self.has_sufficient_balance(amount):
            return False
        self.balance -= amount
        return True

    def transfer(self, amount: float, target: 'Account') -> bool:
        if not self.has_sufficient_balance(amount):
            return False
        self.balance -= amount
        target.deposit(amount)
        return True


class Bank:
    def __init__(self, bank_name: str):
        self.bank_name = bank_name
        self._accounts: list[Account] = []

    def create_new_account(
        self,
        owner_name: str
    ) -> Account:
        new_account = Account(
            id=str(uuid4()),
            owner_name=owner_name,
            balance=0.0
        )
        self._accounts.append(new_account)
        return new_account

    def get_account_by_id(self, account_id: UUID) -> Account | None:
        for account in self._accounts:
            if account.id == account_id:
                return account
        return None
