const API_URL: &str = "https://api.quotable.io/random";

use serde::{Deserialize, Serialize};

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Quote {
    #[serde(rename = "_id")]
    pub id: String,
    pub content: String,
    pub author: String,
    pub tags: Vec<String>,
    pub author_slug: String,
    pub length: i64,
    pub date_added: String,
    pub date_modified: String,
}

#[tokio::main]
async fn main() -> Result<(), reqwest::Error> {
    let quote: Quote = reqwest::get(API_URL).await?.json().await?;
    let content = quote.content.trim();
    let author = quote.author.trim();

    println!("{content}\n{author:>width$}", width = content.len());
    Ok(())
}
