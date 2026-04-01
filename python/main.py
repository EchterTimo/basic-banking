from models import Bank


def main():
    bank = Bank("Sparkasse Berlin")

    acc1 = bank.create_new_account("Max")
    acc2 = bank.create_new_account("Alex")

    print(acc1)
    print(acc2)
    print()

    acc1.deposit(100)
    acc2.deposit(-10)

    print(acc1)
    print(acc2)
    print()

    acc1.transfer(20, acc2)

    print(acc1)
    print(acc2)


if __name__ == '__main__':
    main()
