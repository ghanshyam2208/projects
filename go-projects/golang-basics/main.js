class Deck {
  cards;

  constructor() {
    this.cards = [];
  }

  addCards() {
    this.cards.push("first card");
    this.cards.push("another lol");
    this.cards.push("another card");
  }

  printCards() {
    this.cards.forEach((card, index) => {
      console.log(`Card ${index + 1}: ${card}`);
    });
  }
}

function main() {
  let deck = new Deck();
  deck.addCards();
  deck.printCards();
}

main();
