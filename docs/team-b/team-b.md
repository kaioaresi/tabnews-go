Claro, vou montar um prompt bem direto. Perai.

Aqui vai um prompt copiável, primeiro só para análise. Analise este repositório e proponha o caminho de menor risco e menor esforço para reduzir a quantidade de pipelines geradas no Codefresh. Hoje existe uma lógica que gera pipelines a partir de combinações de squad × environment, criando muitas pipelines com status NEVER_RUN e possivelmente sem diretório correspondente no repo. Quero apenas análise agora, sem alterar arquivos nem excluir nada. Investigue como as pipelines são geradas, onde ocorre a combinação squad × environment e proponha a menor alteração possível para gerar somente pipelines de combinações que realmente existem como diretórios no repositório. Considere como candidatas seguras à remoção apenas as pipelines que estejam NEVER_RUN e não tenham diretório correspondente. Não quero alterar Terraform state, backend, locking, nem refatorar execução ou implementar router/executor agora. Avalie riscos, edge cases, impacto em triggers, approvals e parâmetros. Ao final, traga um plano passo a passo, estratégia de validação e rollback, e uma recomendação GO / NO-GO. Não implemente ainda.

Analise este repositório e proponha o caminho de **menor risco e menor esforço** para reduzir a quantidade de pipelines geradas no Codefresh.

## Contexto

Hoje existe uma lógica que gera pipelines a partir das combinações de `squad × environment`. Isso criou centenas de pipelines, incluindo muitas que nunca foram executadas (`NEVER_RUN`) e que aparentemente não possuem mais um diretório `<squad>/<environment>` correspondente no repositório.

O objetivo imediato é liberar capacidade no Codefresh **sem redesenhar a arquitetura agora**.

A hipótese de quick win é:

1. Parar de gerar o produto cartesiano de squads e environments.
2. Gerar pipelines somente para combinações `<squad>/<environment>` que realmente existem no repositório.
3. Identificar pipelines existentes que atendam simultaneamente aos critérios:
   - status `NEVER_RUN`;
   - não existe diretório `<squad>/<environment>` correspondente.
4. Considerar essas pipelines como candidatas à remoção.

## O que eu quero que você faça

Primeiro, **não altere nenhum arquivo e não exclua nenhuma pipeline**. Faça somente a análise.

Investigue:

- Onde e como as pipelines são geradas atualmente.
- Onde ocorre a combinação `squad × environment`.
- Como determinar programaticamente quais combinações realmente existem no repositório.
- Se é possível substituir a lógica atual por uma geração baseada apenas nos diretórios existentes, com uma alteração pequena e segura.
- Quais pipelines existentes seriam preservadas.
- Quais seriam candidatas à remoção usando o critério `NEVER_RUN + diretório inexistente`.
- Se existe algum caso em que esse critério possa gerar falso positivo.
- Se triggers, approvals, parâmetros, histórico ou qualquer outra configuração das pipelines utilizadas pode ser afetada pela mudança.

## Restrições

Neste quick win, **não quero**:

- alterar Terraform state;
- alterar backend ou locking;
- refatorar a execução do Terraform;
- implementar router/executor genérico;
- redesenhar toda a arquitetura de pipelines;
- alterar pipelines que possuem diretório válido e estão em uso.

O foco é exclusivamente eliminar geração desnecessária e identificar pipelines órfãs com segurança.

## Resultado esperado

Ao final da análise, apresente:

1. Como funciona a geração atual.
2. A causa exata da geração excessiva.
3. A menor alteração necessária para gerar somente pipelines válidas.
4. Quantas pipelines seriam preservadas e quantas poderiam ser removidas, se os dados disponíveis permitirem calcular.
5. Riscos e edge cases encontrados.
6. Como validar a mudança antes da remoção.
7. Estratégia simples de rollback.
8. Arquivos e trechos que precisariam ser alterados.
9. Sua recomendação final: **GO / NO-GO para esse quick win**, com justificativa.

**Não implemente nada ainda.** Quero revisar a análise e o plano antes de autorizar qualquer alteração.
