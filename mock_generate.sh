# package: use_case
mockgen -source use_case/data_reader.go -package use_case -destination use_case/mock_data_reader.go  
mockgen -source use_case/data_writer.go -package use_case -destination use_case/mock_data_writer.go
mockgen -source use_case/field_transformer.go -package use_case -destination use_case/mock_field_transformer.go
