
# Guide to Contributing

The Algorand docs repository hosts Algorand’s core developer documentation. This content is served directly to the Algorand Developer Portal under the Docs section located here: [Algorand Developer Docs](https://developer.algorand.org/docs/). 

## Who We Are

The Algorand technical documentation repository is managed by the Algorand Docs Team which is composed of members from the Developer Relations, Engineering, and Product teams at Algorand Inc.

<a id="charter"></a>**Our Charter**

We work closely and tirelessly with the core Algorand Development Team to produce and present accurate, comprehensive, and up-to-date technical documentation to support the best developer experience on Algorand.

<a id="tenets"></a>**Our Tenets**

* Provide all the building blocks of technical information necessary for developers to build applications or run nodes on Algorand.
* Be the most objective source of technical information. 
* Write clear prose and code to swiftly deliver new concepts to all developers.

## Types of Contributions
As members of an open source project we welcome community contributions in the following ways and we thank you in advance for your help!

### Fixing Errors or Inaccuracies
The Algorand protocol and SDKs are constantly updated with new features that we do our best to stay ahead of, but sometimes portions of our documentation may become out of date. Sometimes we just make mistakes. If you come across typos, code errors, or information that is inaccurate, we encourage you to [log an issue](#how-to-file-an-issue) or correct it through a [pull request](#how-to-create-a-pull-request).

[**Issue Template**](https://github.com/algorand/docs/issues/new?assignees=&labels=minor+update&template=fixing-errors-or-inaccuracies.md&title=)

### Small Improvements for Clarity of Message
We recognize that there may be a clearer way to explain a concept or write a snippet of code and we welcome contributions that can make these improvements on our behalf. 

An improvement can sometimes be subjective, so it helps to include more objective data points such as having interest from a diverse representation of the developer community in the form of upvotes on the related issue or PR.  

[**Issue Template**](https://github.com/algorand/docs/issues/new?assignees=&labels=&template=small-improvements-for-clarity-of-message.md&title=)

### Large Additions or Modifications
The Algorand technical documentation is intimately tied to updates in [go-algorand](https://github.com/algorand/go-algorand), our SDKs, and other open source developer tools found [here](https://github.com/algorand). It may be difficult to have all the context around the developer experience as a whole, so we strongly encourage you to [log an issue](#how-to-file-an-issue) _before_ creating a pull request so we can discuss and let the community weigh in. Similar to the above section, objective data points like issue upvoting from a significant and diverse representation of the developer community will help determine the priority-level for these proposals. 

[**Issue Template**](https://github.com/algorand/docs/issues/new?assignees=&labels=&template=large-additions-or-modifications.md&title=)

#### Tutorial and Solution Submissions
If you have a step-by-step guide for developers or sample application code, we encourage you to [submit it](https://developer.algorand.org/pages/submissions/) as a [Tutorial](https://developer.algorand.org/tutorials/) or [Solution](https://developer.algorand.org/solutions/), respectively, on the [Algorand Developer Portal](https://developer.algorand.org/). Tutorials and Solutions can be more personalized and do not follow the same [charter](#charter) and [tenets](#tenets) as the documentation. View the Style Guide for Tutorials and Solutions [here](https://developer.algorand.org/pages/style-guide).

### Community Projects
Finally, if you developed a tool or application that would benefit members of the Algorand community, feel free to submit a PR updating the [Community Projects](./docs/community.md) page. Include the name of your project, a logo, and a description under the appropriate Project category.  If your category does not exist, please suggest a header and we can update as needed through the PR process. 

The Community page is displayed [here](https://developer.algorand.org/docs/community/) in production.

## How to File an Issue
First search [open issues](https://github.com/algorand/docs/issues) to make sure the issue has not been filed already. If it has not, open a new issue using one of the following three templates or Other, if your request does not fall into any of these categories:

1. [Fixing Errors and Inaccuracies](https://github.com/algorand/docs/issues/new?assignees=&labels=minor+update&template=fixing-errors-or-inaccuracies.md&title=)
2. [Small Improvements for Clarity of Message](https://github.com/algorand/docs/issues/new?assignees=&labels=&template=large-additions-or-modifications.md&title=)
3. [Large Additions or Modifications](https://github.com/algorand/docs/issues/new?assignees=&labels=&template=large-additions-or-modifications.md&title=)


## How to Create a Pull Request

First read the [Types of Contributions](#types-of-contributions) section to make sure a pull request is the best option for your contribution. If so, fork this repo, and initiate your pull request against the `staging` branch following the general process described [here](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork). Make sure to [test your changes](#testing-changes) before submitting. After a pull request is submitted, the docs team will review it and communicate through the comments section of the PR. 

Approved changes will be merged into `staging` and periodically merged into `master` at which point it will be visible on the Algorand Developer Portal. 

### Testing Changes

We use [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/) as the static site generator with several activated plugins that can be found in the [mkdocs.yml](./mkdocs.yml) file.

Run the documentation locally in a web browser to test how your changes will render. The local theme is an out-of-the-box Material theme, which is not exactly what it will look like on the website where we apply our own theme to match the rest of the Developer Portal. However, it should be close enough to determine whether or not your changes will result in formatting errors.

To run locally:

From within the root `docs` directory.

Make sure you have all dependencies installed:

```bash
$ pip3 install -r requirements.txt
```

Run mkdocs and serve the content in your local browser:

```bash
$ mkdocs serve
```

View the docs in your web browser: [http://localhost:8000/](http://localhost:8000/)



## Code of Conduct
 
 See [here](./CODE_OF_CONDUCT.md).


