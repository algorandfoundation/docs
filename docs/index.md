{% load badge_tags %}

title: Algorand Developer Docs

<div class="docs-homepage__challenge-box">
    <div class="challenge-overview-icon"></div>

    <div>
      <p class="button--large docs-homepage__challenge-box__title">Complete Challenges and Earn Badges</p>

      <p>Become a master Algorand Developer by completing coding challenges and getting rewarded with on-chain Badges along the way!</p>
    </div>

    <a href="{% url 'challenges-list' %}" class="button--small button--primary docs-homepage__challenge-box__link">
      View Challenges

      {% include "button-right-arrow.svg" %}
    </a>
</div>

# Top-level sections
<ul class="docs-homepage__card-list">
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-started/basics/what_is_blockchain/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--purple">
                <span class="create-smart-contract-overview-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">Get started</p>
            <p class="typography--regular-body docs-homepage__card__description">Start here to learn the basics</p>
        </a>
    </li>
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--purple">
                <span class="stateless-smart-contracts-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">Get details</p>
            <p class="typography--regular-body docs-homepage__card__description">In-depth guides and explanations of all features</p>
        </a>
    </li>
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./run-a-node/setup/types/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--purple">
                <span class="stateful-smart-contracts-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">Run a node</p>
            <p class="typography--regular-body docs-homepage__card__description">Setup a node and participate in consensus</p>
        </a>
    </li>
</ul>

# Spotlight

<ul class="docs-homepage__card-list">
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/dapps/avm/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--turquoise">
                <span class="payment-with-algos-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">The AVM</p>
            <p class="typography--regular-body docs-homepage__card__description">The Algorand Virtual Machine</p>
        </a>
    </li>
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/dapps/pyteal/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--turquoise">
                <span class="payment-with-algos-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">PyTeal</p>
            <p class="typography--regular-body docs-homepage__card__description">Write smart contracts with Python</p>
        </a>
    </li>
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/walletconnect/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--turquoise">
                <span class="payment-with-algos-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">WalletConnect</p>
            <p class="typography--regular-body docs-homepage__card__description">Integrate mobile wallets with dApps</p>
        </a>
    </li>
        <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/accounts/rekey/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--turquoise">
                <span class="payment-with-algos-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">Rekeying</p>
            <p class="typography--regular-body docs-homepage__card__description">Change the spending key for any account</p>
        </a>
    </li>
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/asa/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--turquoise">
                <span class="payment-with-algos-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">ASAs</p>
            <p class="typography--regular-body docs-homepage__card__description">Tokenize FTs, NFTs, and more with a built-in standard</p>
        </a>
    </li>
    <li class="docs-homepage__card-list-item">
        <a class="docs-homepage__card" href="./get-details/atomic_transfers/">
            <div class="docs-homepage__card__icon-container docs-homepage__card__icon-container--turquoise">
                <span class="payment-with-algos-icon"></span>
            </div>
            <p class="docs-homepage__card__title text-gray--main typography--large-button">Atomic Transfers</p>
            <p class="typography--regular-body docs-homepage__card__description">Group up to 16 transactions that will either all succeed or all fail</p>
        </a>
    </li>
</ul>

{% endverbatim %}
{% generate_challenge_thumbnail challenge_slug="1" %}
{% verbatim %}
