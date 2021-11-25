## Subdomain: Products

Bounded Context: AddProduct
    Input: Product
    Output: ProductDetails or InvalidProductDetails
    Dependencies: CheckProductIdExist, CheckReviews

    Steps:
    1. Search a product or Click a product from Catalog
    2. Validate if product id exist
    3. Get product by id
    4. Return and load the product