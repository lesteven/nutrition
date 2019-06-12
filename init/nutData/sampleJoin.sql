-- 3 way table join

SELECT long_name, nutrient_name, output_value, output_uom, \
    serving_size, serving_size_uom FROM product \
    INNER JOIN nutrient ON product.ndb_no = nutrient.ndb_no \
    INNER JOIN serving ON product.ndb_no = serving.ndb_no \
    WHERE product.ndb_no = '45127487';
